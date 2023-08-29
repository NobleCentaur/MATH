from PIL import Image
import numpy as np
from multiprocessing.pool import ThreadPool

imageSize = 10000
maxIteration = 10000
step = 2.5 / (imageSize - 1)
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
    for a in range((rowNum), (rowNum + 1)):
        for b in range(imageSize):
            complexNum = complex((-2 + (b * step)), (1.25 - (a * step)))
            mandelbrotNum = mandelbrotTest(complexNum)
            if mandelbrotNum == maxIteration:
                array[a][b] = (0, 0, 0)
            if mandelbrotNum != maxIteration:
                colorValue = gradientStart + (gradientStep * mandelbrotNum)
                array[a][b] = colorValue

# variable with next working row
# one thread that reads next working row and 

if __name__=="__main__":
    with ThreadPool() as pool:
        pool.map(mandelbrotRow, range(imageSize))
        
    array = array.astype('uint8')

    new_image = Image.fromarray(array)
    new_image.save('Image.png')