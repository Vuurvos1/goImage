<img align="right" height="100" src="data/icon.png">

# Image viewer

A Blazingly™ Fast image viewer build using go 🔥

## Suported formats

| File extension        | Read | Write | Comment |
| --------------------- | ---- | ----- | ------- |
| ai                    |      |       |         |
| avif                  |      |       |         |
| b64                   |      |       |         |
| bmp                   | ✔    |       |         |
| bpg                   |      |       |         |
| cur                   |      |       |         |
| cur                   |      |       |         |
| dib                   |      |       |         |
| emf                   |      |       |         |
| eps                   |      |       |         |
| exif                  |      |       |         |
| exr                   |      |       |         |
| fax                   |      |       |         |
| fits                  |      |       |         |
| gif                   |      |       |         |
| hdr                   |      |       |         |
| heic                  |      |       |         |
| heif                  |      |       |         |
| ico                   |      |       |         |
| jfif                  |      |       |         |
| jp2                   |      |       |         |
| jpe                   |      |       |         |
| jpeg                  | ✔    |       |         |
| jpg                   | ✔    |       |         |
| jxl                   |      |       |         |
| mjpeg                 |      |       |         |
| pbm                   |      |       |         |
| pcx                   |      |       |         |
| pgm                   |      |       |         |
| png                   | ✔    |       |         |
| ppm                   |      |       |         |
| psd                   |      |       |         |
| qoi                   |      |       |         |
| svg                   |      |       |         |
| tga                   |      |       |         |
| tif                   | ✔    |       |         |
| tiff                  | ✔    |       |         |
| viff                  |      |       |         |
| vx                    |      |       |         |
| webp                  | ✔    |       |         |
| wmf                   |      |       |         |
| wpg                   |      |       |         |
| xbm                   |      |       |         |
| xpm                   |      |       |         |
| RAW and other formats |      |       |         |
| (base-64 string)      |      |       |         |

## Dev reload mode

```bash
pnpm nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run *.go
```
