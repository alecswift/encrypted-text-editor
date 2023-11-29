import json
import time

#function to write to json file
def write_json(new_data, filename='/home/alec/Desktop/code/osu_projects/encrypted_text_editor/password_microservice/data.json'):
    with open(filename,'r+') as file:
        file_data = json.load(file)
        file_data["user_details"].append(new_data)
        file.seek(0)
        json.dump(file_data, file, indent = 4)

#read file for username and password
while(1): 
    file = open("/home/alec/Desktop/code/osu_projects/encrypted_text_editor/password_microservice/user_password.txt", "r+")
    data = file.readlines()
    if data[0] != "User Data Stored":


        #modify data to not have \n at the end
        username = data[0]
        password = data[1]
        data[0] = username.split("\n")[0]
        data[1] = password.split("\n")[0]


        #store data to json file
        userData = {
            data[0]: data[1]
        }

        write_json(userData)
        
        file.close()

        #write confirmation message
        file = open("/home/alec/Desktop/code/osu_projects/encrypted_text_editor/password_microservice/user_password.txt", "w")
        file.write("User Data Stored")
        file.close()
    
    time.sleep(1)
    file.close() 