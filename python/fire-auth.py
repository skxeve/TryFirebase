import firebase_admin
from firebase_admin import credentials, auth
import sys

args = sys.argv

print(args)

firebase_cred = credentials.Certificate("./secrets/try-firebase-auth_adminsdk.json")
firebase_default_app = firebase_admin.initialize_app(firebase_cred)
print(firebase_default_app.name)

def print_user(user):
    try:
        print("User.uid:{}".format(user.uid))
        print("\tUser.display_name:{}".format(user.display_name))
        print("\tUser.phone_number:{}".format(user.phone_number))
        print("\tUser.email:{}".format(user.email))
        print("\tUser.email_verified:{}".format(user.email_verified))
        print("\tUser.photo_url:{}".format(user.photo_url))
        print("\tUser.disabled:{}".format(user.disabled))
        print("\tUser.password:{}".format(user.password))
    except AttributeError as e:
        print("\tforbidden attribute: " + str(e))

try:
    if args[1] == "create":
        print("create block!!")
        user = auth.create_user(
            email='skxeve+localtest190204@gmail.com',
            email_verified=False, # default false
            phone_number='+15555550100',
            password='secretPassword', # no hashed, over 6 length string.
            display_name='John Doe',
            #photo_url='http://www.example.com/12345678/photo.png',
            disabled=False # default false
        )
        print('Sucessfully created new user: {0}'.format(user.uid))
    if args[1] == "get":
        uid = args[2]
        print("get block!!")
        user = auth.get_user(uid)
        print('Success fetched user data: {0}'.format(user))
        print_user(user)
    if args[1] == "list":
        page = auth.list_users()
        print("page: {}".format(str(page)))
        while page:
            print("page.users: {}".format(str(page.users)))
            for user in page.users:
                print_user(user)
            page = page.get_next_page()
            print("page: {}".format(str(page)))

except IndexError as e:
    print("args index error:" + str(e))


"""
# uid is 1~128 length string
user = auth.get_user(uid)
user = auth.get_user_by_email(uid)
print('Success fetched user data: {0}'.format(user.uid))

user = auth.create_user(
    email='user@example.com',
    email_verified=False, # default false
    phone_number='+15555550100',
    password='secretPassword', # no hashed, over 6 length string.
    display_name='John Doe',
    photo_url='http://www.example.com/12345678/photo.png',
    disabled=False # default false
)
print('Sucessfully created new user: {0}'.format(user.uid))

page = auth.list_users()
while page:
    for user in page.users:
        print("User: {}".format(user.uid))
    page = page.get_next_page()

"""
