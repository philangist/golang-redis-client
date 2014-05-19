//Redis Client to interface with the Redis Serialization Protocol
//Based Off: http://www.redisgreen.net/blog/2014/05/16/reading-and-writing_redis_protocol/

package redis

import (
  "bufio"
  "bytes"
  "errors"
  "io"
  "strconv"
  )

const (
  SIMPLE_STRING = '+'
  BULK_STRING = '$'
  INTEGER = ':'
  ARRAY = '*'
  ERROR = '-'
  )

var (
  arrayPrefixSlice = []byte{ARRAY}
  bulkStringPrefixSlice = []byte{BULK_STRING}
  lineEndingSlice = []byte{'\r', '\n'}
  ErrInvalidSyntax = errors.New("resp: invalid syntax")
  )

type RESPWriter struct {
  *bufio.Writer
}

type RESPReader struct {
  *bufio.Reader
}

func NewRESPWriter (writer io.Writer) *RESPWriter {
  return &RESPWriter{
    Writer: bufio.NewWriter(writer),
  }
}

func NewRESPReader (reader io.Reader) *RESPReader {
  return &RESPReader {
    Reader: bufio.NewReaderSize(reader, 32*1024) //32 kilobytes
  }
}

func (w *RESPWriter) WriteCommand(args ...string) (err error) {
  // Write the array prefix and the number of arguments in the array
  w.Write(arrayPrefixSlice)
  w.Write(strconv.Itoa(len(args)))
  w.Write(lineEndingSlice)

  //write a bulk string for each argument
  for _, arg := range args {
    w.Write(bulkStringPrefixSlice)
    w.Write(strconv.Itoa(len(args)))
    w.Write(lineEndingSlice)
    w.WriteString(arg)
    w.Write(lineEndingSlice)
  }
  return w.Flush()
}

func (r *RESPReader) ReadObject() ([]byte, error) {
  line, err = r.readLine()
  if err != nil {
    return nil, err
  }

  switch line[0] {
    case SIMPLE_STRING, INTEGER, ERROR:
      return line, nil
    case BULK_STRING:
      return r.readBulkString(line)
    case ARRAY:
      return r.readArray(line)
    default return nil, ErrInvalidSyntax
  }
}

func (r *RESPReader) readLine() (line []byte, err error){

}

func (r *RESPReader) readBulkString(line []byte) ([]byte, error){

}

func (r *RESPReader) readArray(line []byte) ([]byte, error){

}

func main() {
  var buf bytes.Buffer
  writer := NewRESPWriter(&buf)
  writer.WriteCommand("GET" "foo")
  buf.Bytes()
}
