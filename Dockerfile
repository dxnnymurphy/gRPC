FROM python:3.8
ENV http_proxy http://10.150.32.10:3128
ENV https_proxy http://10.150.32.10:3128
WORKDIR /py
COPY py/requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["python", "py/python_server.py"]