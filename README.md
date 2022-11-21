# MapBuilder Web Service
The web service to generate SVG vector shapes of countries from Natural Earth database: https://www.naturalearthdata.com/downloads/.

# Build

To build the application, you need to have Go installed: https://go.dev/dl/ and run the following command:

go build cmd/mapbuilder/mapbuilder

It will create the `mapbuilder` executable binary file.

# Run

Execute command:

```
./mapbuilder -p <port-number>
```

Before running this command, ensure that `geodata` folder with data exists in a folder with `mapbuilder` file.

It will run a web server that will listen for SVG generation requests on specified `port-number`. If port number not specified, then it will listen on port `6001`.

# Use

MapBuilder web service exposes the following endpoint:

`POST /map`

That requires POST body in JSON format:

```javascript
{
  "countries": ["PER","NOR"],       // Array of country ISO codes
  "scale": "10m",                   // Scale of geodata. 
                                    // Should be one of 10m, 50m or 110m
  "width": "width",                 // Width o resulting SVG file in pixels
  "height": "height"                // Height of resulting SVG file in pixels 
}
```

`countries` argument must include only country codes that exist in `geodata` folder under specified scale.

If request formed correctly, web service will respond with SVG file body, compressed by gzip.

# Example

The client application example, that uses this service available here:

https://maps.germanov.dev
