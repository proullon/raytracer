window.onload = function() {
	var size = 900;
	var canvas = document.getElementById('canvas');
	var ctx = canvas.getContext('2d');
	var w = size;
	var h = size;

	canvas.width  = w;
	canvas.height = h;

    this.socket = new WebSocket('ws://' + window.location.host + '/ws/raytracer');
    this.socket.onmessage = function (event) {
        // console.log("socket.onmessage : " + event)
	    var pixels = JSON.parse(event.data);
	    i = 0;
	    while (pixels[i] != null) {
		    drawPixel(ctx, pixels[i]);
		    i++;	    	
	    }
    }
    this.socket.onclose = function() {
        console.log("socket.onclose");
    }

	// image = getImage();
	// draw(ctx, image, w, h);
}


function getImage(size) {

	var xhr = getXMLHttpRequest();

	xhr.open("GET", '/api/raytracer', false);
	xhr.send(null);
	if(xhr.readyState != 4 || (xhr.status != 200 && xhr.status != 0)) // Code == 0 en local
		throw new Error("Cannot fetch image from server (HTTP : " + xhr.status + ").");
	var imgJsonData = xhr.responseText;

	var img = JSON.parse(imgJsonData);

	return img;	
}

function drawPixel(ctx, pixel) {
	var imageData = ctx.createImageData(1, 1)

	imageData.data[0] = pixel.R
	imageData.data[1] = pixel.G
	imageData.data[2] = pixel.B
	imageData.data[3] = 255

	ctx.putImageData(imageData, pixel.X, pixel.Y)
}

function drawBackground(ctx, w, h) {
	var imageData = ctx.createImageData(w, h)

	var index = 0;
	var pixels = w * h * 4
	while (index < pixels) {
		imageData.data[index]    = 0
		imageData.data[index +1] = 0
		imageData.data[index +2] = 0
		imageData.data[index +3] = 255		
		index += 4;
	}

	ctx.putImageData(imageData, 0, 0)
}

function draw(ctx, img, w, h) {
	var imageData = ctx.createImageData(1, 1)

	var x = 0;
	while (x < h) {

		var y = 0;
		while (y < w) {

			color = img[x][y];
			console.log("R:" + color.R + " G:" + color.G + " B:" + color.B);
			imageData.data[0] = color.R;
			imageData.data[1] = color.G;
			imageData.data[2] = color.B;
			imageData.data[3] = 255;
			ctx.putImageData(imageData, x, y)

			y++;
		}

		x++;
	}
}