from PIL import Image
import numpy as np
from multiprocessing.pool import ThreadPool
from timeit import default_timer as timer

# Determines the image pixel width and length
imageSize = int(input("Render Size : "))
imageSharpness = int(input(f"Render Sharpness [1 - {imageSize}]: "))
# Determines how precise the gradient is
maxIteration = imageSize / imageSharpness
# Default should be 2.5
valueRange = 2.5
# Default should be (-2, 1.25)
startingPoint = (-2, 1.25)

step = valueRange / (imageSize - 1)
gradientStart = (255, 0, 0)
gradientEnd = (5, 0, 0)
gradientStep = np.zeros((1, 3), dtype=np.float32)
gradientStep = (np.subtract(gradientEnd, gradientStart))

array = np.zeros((imageSize, imageSize, 3), dtype=np.float32)
for x in range(0, imageSize):
    for i in range(0, imageSize):
        array[x][i] = [255, 255, 255]

def mandelbrotTest(C): 
    Z = 0
    n = 0
    while abs(Z) <= 2 and n < maxIteration:
        Z = Z * Z + C
        n += 1
    return(n)

def mandelbrotRow(rowNum):
    for b in range(imageSize):
        complexNum = complex((startingPoint[0] + (b * step)), (startingPoint[1] - (rowNum * step)))
        mandelbrotNum = mandelbrotTest(complexNum)
        if mandelbrotNum == maxIteration:
            array[rowNum][b] = (0, 0, 0)
        if mandelbrotNum != maxIteration:
            colorValue = gradientStart + (gradientStep * mandelbrotNum)
            array[rowNum][b] = colorValue

# variable with next working row
# one thread that reads next working row and 

if __name__=="__main__":
    startTime = timer()
    np.round(startTime, 3)

    with ThreadPool() as pool:
        pool.map(mandelbrotRow, range(imageSize))
        
    array = array.astype('uint8')
    new_image = Image.fromarray(array)
    new_image.save('Mandelbrot.png')
    
    endTime = timer()
    np.round(endTime, 3)
    print("")
    print(f"Finished! Took {endTime - startTime} seconds.")
    print("")