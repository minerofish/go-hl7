# go-astm ![build status](https://travis-ci.org/78bit/uuid.svg?branch=master)

Golang library for handling hl7 2.x Procotol

###### Install
`go get github.com/DRK-Blutspende-BaWueHe/go-hl7`

## Features
  - Encoding 
    - UTF8 
    - ASCII
    - Windows1250 
    - Windows1251 
    - Windows1252 
    - DOS852 
    - DOS855 
    - DOS866 
    - ISO8859_1
  - Timezone Support
  - Marshal/Unmarshal function

## Quick Start

The following Go code decodes a hl7 read from a File.

``` go
fileData, err := ioutil.ReadFile("tbd.hl7")
if err != nil {
  log.Fatal(err)		
}

var message 
err := hl7.Unmarshal(
		[]byte(filedata),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)
    
if err != nil {
   log.Fatal(err)		
}
```

To identify the version and type of the message, you can use the IdentifyMessage function.

```go
	var data string
	messageType, protocolVersion, err := hl7.IdentifyMessage(
		[]byte(data),
		hl7.EncodingUTF8,
	)
```
