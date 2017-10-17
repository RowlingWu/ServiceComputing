package main

import (
    "os"
    "flag"
    "syscall"
    "os/exec"
    "bufio"
    "fmt"
)

const BUFSIZ = 16 * 1024
const MAX_INT = int(^uint(0) >> 1)

type selpg_args struct {
    start_page, end_page int
    in_filename string
    page_len int
    page_type bool // False: '-l'. True: '-f'
    print_dest string
}

var progname string /* programe name, for error message */

func process_args(psa *selpg_args) {
    flag.IntVar(&psa.start_page, "s", -1, "start page (>= 1)")
    flag.IntVar(&psa.end_page, "e", -1, "end page (>= start page)")
    flag.IntVar(&psa.page_len, "l", -1, "define lines per page (default 72)")
    flag.BoolVar(&psa.page_type, "f", false, "False: -l. True: -f")
    flag.StringVar(&psa.print_dest, "d", "", "the result will be written in this file (default is stdout)")
    flag.Parse()
    flag.Usage = func() {
        os.Stdout.WriteString("\nUSAGE: " + progname + " -s start_page -e end_page [ -f | -l lines_per_page ] [ -d dest ] [ in_filename ]\n")
    }
    /* check whether the input is legal */
    if psa.start_page < 1 || psa.start_page >= MAX_INT {
        os.Stderr.WriteString(progname + ": invalid start page\n")
        flag.Usage()
        os.Exit(1)
    }

    if psa.end_page < 1 || psa.end_page >= MAX_INT || psa.end_page < psa.start_page {
        os.Stderr.WriteString(progname + ": invalid end page\n")
        flag.Usage()
        os.Exit(2)
    }

    if psa.page_type {  // page type : -f
        if psa.page_len != -1 {
            os.Stderr.WriteString(progname + ": -f conflicts with -l\n")
            flag.Usage()
            os.Exit(3)
        }
    } else {   // page type : -l
        switch {
        case psa.page_len == -1:
            psa.page_len = 72  // default 72 lines per page
        case psa.page_len < 0:
            os.Stderr.WriteString(progname + ": invalid page length\n")
            flag.Usage()
            os.Exit(4)
        }
    }

    if flag.NArg() > 1 {
        os.Stderr.WriteString(progname + ": unknown options:\n")
        for _, opt := range flag.Args() {
            os.Stderr.WriteString(opt + "\n")
        }
        flag.Usage()
        os.Exit(6)
    }

    if flag.NArg() == 1 {
        psa.in_filename = flag.Args()[0]
    }

    if len(psa.in_filename) != 0 {
        _, err := os.Stat(psa.in_filename)
        if err != nil || os.IsNotExist(err) {
            os.Stderr.WriteString(progname + ": input file \"" + psa.in_filename + "\" does not exist\n")
            os.Exit(7)
        }
        /* check if file is readable */
        err = syscall.Access(psa.in_filename, syscall.O_RDONLY)
        if err != nil {
            os.Stderr.WriteString(progname + ": input file \"" + psa.in_filename + "\" exists but cannot be read\n")
            os.Exit(8)
        }
    }
}

func process_input(sa selpg_args) {
    var (
        line_ctr int
        cur_page int
        err error
    )
    /* set the input source */
    fin := os.Stdin
    if len(sa.in_filename) != 0 {
        fin, err = os.Open(sa.in_filename)
        if err != nil {
            os.Stderr.WriteString(progname + ": could not open input file \"" + sa.in_filename + "\"\n")
            os.Exit(12)
        } else {
            defer fin.Close()
        }
    }

    /* set the output destination */
    fout := os.Stdout
    if len(sa.print_dest) != 0 {
        fout, err = os.Create(sa.print_dest)
        if err != nil {
            os.Stderr.WriteString(progname + ": " + err.Error() + "\n")
            os.Exit(13)
        } else {
            defer fout.Close()
        }
    }

    cmd := exec.Command("cat", "-n")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        os.Stdout.Write([]byte(err.Error() + "\n"))
        os.Exit(1)
    }

    line := bufio.NewScanner(fin)
    cur_page = 1
    if !sa.page_type {
        line_ctr = 0
        for line.Scan() && cur_page <= sa.end_page {
            if cur_page >= sa.start_page {
                stdin.Write([]byte(line.Text() + "\n"))
            }
            line_ctr += 1
            if line_ctr == sa.page_len {
                line_ctr = 0
                cur_page += 1
            }
        }
    } else {
        for line.Scan() && cur_page <= sa.end_page {
            for _, c := range line.Text() {
                if c == '\f' {
                    cur_page += 1
                }
                if cur_page >= sa.start_page {
                    stdin.Write([]byte(string(c)))
                }
            }
            stdin.Write([]byte("\n"))
        }
    }
    if cur_page < sa.end_page {
        fmt.Println("Reached EOF before reaching -e end_page")
    }
    cmd.Stdout = fout
    stdin.Close()
    cmd.Start()
}

func main() {
    var sa selpg_args
    progname = "selpg"

    process_args(&sa)
    process_input(sa)
}
