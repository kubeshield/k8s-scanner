# Use an official Python runtime as the base image
FROM python:3.8-slim-buster

# Set the working directory in the container
WORKDIR /app

# Copy the requirements file into the container
COPY requirements.txt .

# Install any needed packages specified in requirements.txt
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Copy the rest of the application code into the container
COPY . .

# Make port 80 available to the world outside the container
EXPOSE 80

# Define environment variable
ENV NAME World

# Run the command to start the application when the container launches
CMD ["python", "app.py"]

