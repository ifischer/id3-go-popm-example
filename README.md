# id3-go-popm-example

Example how to read popularimeter frame with id3-go.

## Usage

```
go build -o popm-read
./popm-read samples/with_popm.mp3
./popm-read samples/without_popm.mp3
```

## About samples

`samples` directory contains one file with and one without popularimeter.

Mutagen-inspect output:

```
-- samples/without_popm.mp3
- MPEG 1 layer 3, 320000 bps (CBR, LAME 3.98.1+), 44100 Hz, 2 chn, 0.00 seconds (audio/mp3)
TALB=Aleefee
TBPM=123
TCON=House
TDRC=2011-01-01
TENC=Lame 3.97
TIT2=AleeFee
TKEY=5d
TLAN=eng
TLEN=472601
TPE1=Andhim
TPUB=sorted
TRCK=1/3
TSSE=Lavf57.83.100
TXXX=Catalog #=SUNSETHAND003
TXXX=Release type=Normal
TXXX=Rip date=2011-05-12
TXXX=Ripping tool=EAC
TXXX=Source=WEB
TXXX=Supplier=Lame 3.97
TXXX=comment=#openair

-- samples/with_popm.mp3
- MPEG 1 layer 3, 320000 bps (CBR, LAME 3.98.1+), 44100 Hz, 2 chn, 0.00 seconds (audio/mp3)
POPM=traktor@native-instruments.de=0 255/255
TALB=Aleefee
TBPM=123
TCON=House
TDRC=2011-01-01
TENC=Lame 3.97
TIT2=AleeFee
TKEY=5d
TLAN=eng
TLEN=472601
TPE1=Andhim
TPUB=sorted
TRCK=1/3
TSSE=Lavf57.83.100
TXXX=Catalog #=SUNSETHAND003
TXXX=Release type=Normal
TXXX=Rip date=2011-05-12
TXXX=Ripping tool=EAC
TXXX=Source=WEB
TXXX=Supplier=Lame 3.97
TXXX=comment=#openair
```


