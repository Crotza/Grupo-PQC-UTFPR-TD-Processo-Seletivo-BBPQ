from PIL import Image
import numpy as np
import matplotlib.pyplot as plt
import pytesseract

encrypted_image_path = 'encrypted1.png'
random_image_path = 'random-image.png'

encrypted_image = Image.open(encrypted_image_path)
random_image = Image.open(random_image_path)

encrypted_image_array = np.array(encrypted_image)
random_image_array = np.array(random_image)

# def add(image1, image2):
#     return np.clip(image1 + image2, 0, 255)             # Clipped at 255 to avoid overflow

# def subtract(image1, image2):
#     return np.clip(image1 - image2, 0, 255)

# def multiply(image1, image2):
#     result = (image1.astype(np.float32) * image2.astype(np.float32)) / 255.0
#     return np.clip(result, 0, 255).astype(np.uint8)

# def divide(image1, image2):
#     result = np.divide(image1, image2 + 1e-9) * 255     # Avoid division by zero
#     return np.clip(result, 0, 255).astype(np.uint8)

# def diff(image1, image2):
#     return np.abs(image1 - image2)

# def mod(image1, image2):
#     safe_array2 = np.where(image2 == 0, 1, image2)      # 0 -> 1 to avoid division by zero
#     return np.mod(image1, safe_array2)

# def and_images(image1, image2):
#     return np.bitwise_and(image1, image2)

# def or_images(image1, image2):
#     return np.bitwise_or(image1, image2)

def xor_images(image1, image2):
    return np.bitwise_xor(image1, image2)

if encrypted_image_array.shape == random_image_array.shape:

    # add_result = add(encrypted_image_array, random_image_array)
    # subtract_result = subtract(encrypted_image_array, random_image_array)
    # multiply_result = multiply(encrypted_image_array, random_image_array)
    # divide_result = divide(encrypted_image_array, random_image_array)
    # diff_result = diff(encrypted_image_array, random_image_array)
    # mod_result = mod(encrypted_image_array, random_image_array)
    # and_result = and_images(encrypted_image_array, random_image_array)
    # or_result = or_images(encrypted_image_array, random_image_array)
    xor_result = xor_images(encrypted_image_array, random_image_array)

    # add_image = Image.fromarray(add_result)
    # subtract_image = Image.fromarray(subtract_result)
    # multiply_image = Image.fromarray(multiply_result)
    # divide_image = Image.fromarray(divide_result)
    # diff_image = Image.fromarray(diff_result)
    # mod_image = Image.fromarray(mod_result)
    # and_image = Image.fromarray(and_result)
    # or_image = Image.fromarray(or_result)
    xor_image = Image.fromarray(xor_result)

    flag = 'grupocpq-fAJMAPAq5VK7'
    print(f"Considering the image shown after XOR operation, the FLAG is: {flag}")

    plt.imshow(xor_image)
    plt.title("XOR Result")
    plt.axis('off')
    plt.show()

else:
    print("The images do not have the same dimensions, unable to proceed:")
    print(f" Encrypted Image shape: {encrypted_image_array.shape}")
    print(f" Random Image shape: {random_image_array.shape}")