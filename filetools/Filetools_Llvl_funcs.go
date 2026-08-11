package filetools

import (
	"fmt"
	"log"
	"os"
)

func Parser_logerr(logfname string, ctw error) {
	f, sa := os.OpenFile(logfname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0654)
	f.Write([]byte(ctw.Error()))
	f.Close()
	if sa != nil {
		log.Fatal(sa)
		fmt.Printf("Unable to log parser error")
	}

}

func Sizef(fil_name string) int64 {
	sizeds, _ := os.Stat(fil_name)
	var i int64 = sizeds.Size()
	return i
}

func Writerfunc(name_file string, data []byte) int { //writerfunc: simple writer func| return 1 = func ran fine | return 0 = there was an error in the write action
	f, sa := os.OpenFile(name_file, os.O_CREATE|os.O_WRONLY, 0654)
	f.Write(data)
	if sa == nil {
		f.Close()
		return 1
	} else {
		return 0
	}

}

func Cspace_b(slice []byte) []byte { //clears spaces | LINE FEEDS from slice, as long as slice is []byte

	for j := 0; j <= (len(slice) - 1); j++ {
		x := slice[j]
		if x != byte('\u0020') && x != byte('\u0010') {
			slice[j] = x

		} else {
			slice = append(slice[:(j)], slice[(j+1):]...)
		}

	}
	return slice
}
func Cspace_s(slice []string) []string { //clears spaces | LINE FEEDS from slice, as long as SLICE is []string

	for j := 0; j <= (len(slice) - 1); j++ {
		x := slice[j]
		if x != string('\u0020') && x != string('\f') {
			slice[j] = x

		} else {
			slice = append(slice[:(j)], slice[(j+1):]...)
		}

	}
	return slice
}
func Nwriterfunc_nspace(name_file string, data []byte, line_size int) int { //Nwriterfunc_nspace: writer func, but adds \n after LINE_SIZE elements from DATA, also uses Cspace to remove ascii/utf-8 spaces| return 1 = func ran fine | return 0 = there was an error in the write action
	f, sa := os.OpenFile(name_file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0654)
	le := len(data)
	if le < line_size {
		f.Write(data)
	} else {
		n_nospace := Cspace_b(data)
		for i := 0; i <= (len(n_nospace) - 1); i += line_size {
			n := n_nospace[i : i+line_size]
			f.Write(n)
			f.Write([]byte("\n"))
		}
	}
	if sa == nil {
		f.Close()
		return 1
	} else {
		return 0
	}

}

func Nwriterfunc(name_file string, data []byte, line_size int) int { //Nwriterfunc: writer func, but adds \n after LINE_SIZE elements from DATA| return 1 = func ran fine | return 0 = there was an error in the write action
	f, sa := os.OpenFile(name_file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0654)
	le := len(data)
	if le < line_size {
		f.Write(data)
	} else {
		n_nospace := data
		for i := 0; i <= (len(n_nospace) - 1); i += line_size {
			n := n_nospace[i : i+line_size]
			f.Write(n)
			f.Write([]byte("\n"))
		}
	}
	if sa == nil {
		f.Close()
		return 1
	} else {
		return 0
	}

}

func File_read(fname string, p_log_path string) *File_whole {
	y := Sizef(fname)
	f, sa := os.Open(fname)
	var b_beingread = make([]byte, y)
	if sa != nil {
		log.Fatal(sa)
		if p_log_path == "none" {
			fmt.Printf("no log file location has been given")
		} else {
			Parser_logerr(p_log_path, sa)
		}

	} else {

		if _, fread := f.Read(b_beingread); fread != nil {
			log.Fatal(fread)
			Parser_logerr(p_log_path, fread)
		}
		f.Close()

	}
	g := File_whole{File: b_beingread}
	return &g
}

func (f *File_whole) File_cut() *File_cutted { //splits IN as 4 easier-ly*?* OUTS (q_1 to q_4) and remainder of len(in)/4
	var q_1, q_2, q_3, q_4, remainder []byte
	lenof_in := len(f.File)
	switch lenof_in % 4 {
	case 3: //This is for the cases where Lenof_in = n*5; in other words, if lenof_in is a multipĺe of/divisable by 5.
		q_1 = f.File[0:(lenof_in / 4)]
		q_2 = f.File[(len(q_1) + 1):(2 * lenof_in / 4)]
		q_3 = f.File[(len(q_1) + len(q_2) + 2):(3 * lenof_in / 4)]
		q_4 = f.File[(len(q_1) + len(q_2) + len(q_3)):(lenof_in - (lenof_in % 4))]
	case 1: //This is for cases where lenof_in = n*3; a.k.a. lenof_in is a multiple of/divisable by 3
		q_1 = f.File[0:(lenof_in / 4)]
		q_2 = f.File[(len(q_1)):(2 * lenof_in / 4)]
		q_3 = f.File[(len(q_1) + len(q_2)):(3 * lenof_in / 4)]
		q_4 = f.File[(len(q_1) + len(q_2) + len(q_3)):(lenof_in - (lenof_in % 4))]
	}
	remainder = f.File[(len(q_1) + len(q_2) + len(q_3) + len(q_4)):lenof_in]
	retu := File_cutted{File_p1: q_1, File_p2: q_2, File_p3: q_3, File_p4: q_4, File_remain: remainder}
	return &retu
}

//
//func File_restruc_hreadab(b []byte) []string {
//	fstruc_blen := (len(b) / 4)
//	var apost []string
//	for i := 1; i <= fstruc_blen; i++ {
//		if b[i] == 0 {
//			apost[i] = "\n"
//		} else {
//			apost[i] = string(b[i])
//		}
//	}
//	return apost
//}
