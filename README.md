samuneiru (thumbnail) Ver.1.0.2
===============================

これはなに？
------------

画像ファイルをリサイズします。入力ファイルのフォーマットはjpg,png,gifから自動判定します。

使い方
------

```
# 標準入力の画像データを幅500pxにして標準出力へ
$ cat input.png | samuneiru -width 500 > output.png

# ファイル名指定版
$ samuneiru -width 500 -ifile input.png -ofile output.png

# 出力フォーマットをjpgに
$ samuneiru -width 500 -ifile input.png -ofile output.jpg -oformat jpg

# HELP
$ samuneiru -help
samuneiru: The image thumbnailer. (Ver.1.0.2)
usage: ./samuneiru [-width | -height] [other options]
  -height uint
      resize height. (0 is keep aspect ratio)
  -help
      print this help message.
  -ifile string
      input-file name. ('-' is stdin)  (default "-")
  -ofile string
      output-file name. ('-' is stdout)  (default "-")
  -oformat string
      output-file format. [jpg|png|gif] ('-' is same as input) (default "-")
  -width uint
      resize width. (0 is keep aspect ratio)
```

License
-------

MIT

Author
------

Hideyuki Hirauchi &lt;hirauchi@ideapump.net&gt;
