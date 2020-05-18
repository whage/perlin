My experiment with Perlin noise generation.

```
package main

import (
	"github.com/whage/perlin"
)

func main() {
	perlin.CreatePPM(600, 600, 15, 15)
}
```

Wonderful links:
- [Matt Zucker's FAQ](https://mzucker.github.io/html/perlin-noise-math-faq.html)
- [Scratchapixel - Perlin nose part 2](https://www.scratchapixel.com/lessons/procedural-generation-virtual-worlds%20/perlin-noise-part-2?url=procedural-generation-virtual-worlds%20/perlin-noise-part-2)
- [The book of shaders](https://thebookofshaders.com/11/)
