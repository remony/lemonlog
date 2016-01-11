# Lemon Log

A simple Goji v1 middleware to provide logging

# install

``` go get  github.com/remony/lemonlog```

# To use

``` goji.use(lemonlog.logger) ```

> Please note this is provided and personalized towards my Honours project however open to modifications.

- Displays **METHOD**, **Time between request and response** and **Reques URL**
- Current displays only when path == ```/api```

## Possible extensions
- Provide filters instead of just /api