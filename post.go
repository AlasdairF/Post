package post

import (
	"net/http"
	"net/url"
	"github.com/AlasdairF/Custom"
)

type poster struct {
	buf *custom.Buffer
	used bool
}

func New() *poster {
	return &poster{buf: custom.NewBuffer(0)}
}

func (p *poster) Add(key, val string) {
	w := p.buf
	if used {
		w.WriteByte('&')
	} else {
		used = true
	}
	w.WriteString(key)
	w.WriteByte('=')
	w.WriteString(val)
}

func (p *poster) AddEscaped(key, val string) {
	w := p.buf
	if used {
		w.WriteByte('&')
	} else {
		used = true
	}
	w.WriteString(key)
	w.WriteByte('=')
	w.WriteString(url.EncodeQuery(val))
}

func (p *poster) Close() {
	p.buf.Close()
}

func (p *poster) Bytes() []byte {
	return p.buf.Bytes()
}

func (p *poster) POST(uri string) error {
	hc := http.Client{}
    req, err := http.NewRequest(`POST`, uri, p.buf.Reader())
	if err != nil {
		return err
	}
    req.Header.Add(`Content-Type`, `application/x-www-form-urlencoded`)
    _, err = hc.Do(req)
	return err
}
