FROM python:3.9-slim as building
RUN apt-get update
RUN apt-get install -y --no-install-recommends \
build-essential gcc 

WORKDIR /usr/app
RUN python -m venv /usr/app/venv
ENV PATH="/usr/app/venv/bin:$PATH"

COPY requirements.txt .
RUN pip install -r requirements.txt

# ---
FROM python:3.9-slim
WORKDIR /usr/app
COPY --from=building /usr/app/venv ./venv
COPY . .

ENV PATH="/usr/app/venv/bin:$PATH"

EXPOSE 50003

CMD [ "./server.py" ]
