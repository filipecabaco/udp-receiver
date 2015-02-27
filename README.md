# udp-receiver
Small UDP receiver in Golang

VLC:
  MRL: dshow://
  Emission String: :sout=#transcode{acodec=none}:udp{dst=127.0.0.1:1234} :sout-keep
  Note: Still buggy, need to find better config
