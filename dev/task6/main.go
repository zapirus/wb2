package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Cut struct {
	sl     []string
	Fields []string
	Del    string
	Sep    bool
	Total  string
}

func (c *Cut) split(text string) []string {
	return strings.Split(text, c.Del)
}

func (c *Cut) Cut(text string) string {

	c.sl = c.split(text)
	if len(c.sl) <= 1 {
		if c.Sep {
			return ""
		}
		c.Total = c.sl[0]
		return c.Total
	}

	for _, v := range c.Fields {
		j, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}

		if len(c.sl)-1 > j {
			j -= 1
			if j < 0 {
				j = 0
			}
			c.Total += c.sl[j] + " "
		}
	}

	return c.Total
}

var fields = flag.String("f", "", "выбрать поля (колонки)")
var del = flag.String("d", "\t", "использовать другой разделитель")
var sep = flag.Bool("s", false, "только строки с разделителем")

func main() {
	flag.Parse()
	text := flag.Arg(0)

	c := Cut{
		Fields: strings.Split(*fields, ","),
		Del:    *del,
		Sep:    *sep,
	}

	res := c.Cut(text)
	fmt.Println(res)
}
