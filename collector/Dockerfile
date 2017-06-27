FROM python:2.7

WORKDIR /collector
ADD . .

RUN pip install -r requirements.txt

EXPOSE 8000 

CMD ["python", "main.py"]
