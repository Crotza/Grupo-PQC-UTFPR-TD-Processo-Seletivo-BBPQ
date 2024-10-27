FROM python:3.9-slim

RUN apt-get update && apt-get install -y \
    tesseract-ocr \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY random-image.png .
COPY encrypted1.png .
COPY flag-encrypted.py .

CMD ["python", "flag-encrypted.py"]