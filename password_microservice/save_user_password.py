import json
import time

DATA_FILE_PATH = "password_microservice/data.json"
COMM_FILE_PATH = "password_microservice/user_password.txt"

def main():
    while(1): 
        read_file = open(COMM_FILE_PATH, "r+")
        data = read_file.readlines()
        if data[0] != "User Data Stored":
            write_user_password(data)
            read_file.close()
            write_conf_message(read_file)

        time.sleep(1)
        read_file.close()

def write_user_password(data):
    #modify data to not have \n at the end
    username, password = data[0], data[1] 
    data[0], data[1] = username.split("\n")[0], password.split("\n")[0]

    #store data to json file
    userData = {
        data[0]: data[1]
    }

    write_json(userData)

def write_conf_message(read_file):
    read_file = open(COMM_FILE_PATH, "w")
    read_file.write("User Data Stored")
    read_file.close()

def write_json(new_data, filename=DATA_FILE_PATH):
    with open(filename,'r+') as file:
        file_data = json.load(file)
        file_data["user_details"].append(new_data)
        file.seek(0)
        json.dump(file_data, file, indent = 4)

if __name__ == "__main__":
    main()