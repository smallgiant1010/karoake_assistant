FROM python:3.11-slim

WORKDIR /app

RUN pip install pytest requests

COPY . .

CMD ["pytest", "-v", "--tb=short"]