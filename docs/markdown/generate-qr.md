#### Data

`data` can be any string with the length less than 2953 bytes, it's the maximum value that QR code can store

##### Colors

`backgroundColor` and `foregroundColor` are the hex RGB representation. The length of the color is either 3 or 6.  
Valid examples: `#fff`, `#1f1f1f`

#### Size

`size` controls size of the image with QR code, not the actual QR code.

#### Border size

`borderSize` is the space between the edge of an image and the edge of a QR-code.
> Note: the bigger border size, fewer space left for the actual QR code, so it'll appear smaller

#### Logo

Logo in the center of QR code is controlled by `logo` and `logoScale` fields.  
`logo` can be either base64 encoded image or URL to the image. Valid image types: PNG or JPEG.

`logoScale` controls how big logo will be relatively to the QR code size (not the image, but resized QR code, if borders
applied).  
Logo can take up to 25% of the QR code. Hence, the maxiumum value is `0.25`

#### Recovery Levels

Recovery Levels control how much data will be used to duplicate data.
> Note: with higher recovery level, you get more chance that QR code will be scanned even if corrupted.