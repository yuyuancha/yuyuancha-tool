## 安裝步驟

> 步驟為 MacOS 環境執行。

- 安裝 `opencv` 套件。
```shell
brew install opencv
```

- 安裝 `pkgconfig` 套件。
```shell
brew install pkgconfig
```

- 安裝 `gocv`。
```shell
go get -u -d gocv.io/x/gocv
```
## 問題排除

### 執行 `tessbrige` 圖像內容辨識時，噴出的錯誤訊息。

```
tessbridge.cpp:5:10: fatal error: 'leptonica/allheaders.h' file not found
```

解決方法：

```shell
export LIBRARY_PATH="/opt/homebrew/lib"
export CPATH="/opt/homebrew/include"
source ~/.zshrc
```

> 參考網址：https://github.com/otiai10/gosseract/issues/234
