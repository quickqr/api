### Data

`data` can be any string with the length less than 2953 bytes, it's the maximum value that QR code can store

#### Colors

`backgroundColor` and `foregroundColor` are the hex RGB representation. The length of the color is either 3 or 6.  
Valid examples: `#fff`, `#1f1f1f`

### Size

`size` controls size of the image with QR code, not the actual QR code.

### Quiet Zone

`quietZone` is the space between the edge of an image and the edge of a QR-code.
> Note: with bigger quiet zone, fewer space left for the actual QR code, so it'll appear smaller

### Styling

You can style QR code via `finder`, `module` and `gap` values. See reference below

### Gradients

Gradient is set up via `gradientDirection` and `gradientColors` variables (see more in docs below)

### Logo

Logo in the center of QR code is controlled by `logo` and `logoScale` fields.  
`logo` can be either base64 encoded image or URL to the image. Valid image types: PNG or JPEG.

`logoSpace` adds space around logo, QR code will look cleaner

### Recovery Levels

Recovery Levels control how much data will be used to duplicate data.
> Note: with higher recovery level, you get more chance that QR code will be scanned even if corrupted.

### Version

You can force version (and max capacity, then) with `version.`
> Be aware of errors if data overflows max capacity of supplied version