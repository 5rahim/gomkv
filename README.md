# gomkv

[Matroska](https://www.matroska.org) parser for Go. Fork of [github.com/remko/go-mkvparse](https://github.com/remko/go-mkvparse).

Features:

- Supports [all Matroska elements](https://www.matroska.org/technical/specs/index.html)
- Supports short-circuiting the parser, making it possible to 
read specific data (e.g. title, author) without reading the
entire file (see the `mkvtags` example)
- Also works with [WebM](https://www.webmproject.org) (`.webm`) files
- Supports streaming
- Event-based push API
- No dependencies
