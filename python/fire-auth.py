import firebase_admin
from firebase_admin import credentials, auth

firebase_cred = credentials.Certificate("./secrets/try-firebase-auth_adminsdk.json")
firebase_default_app = firebase_admin.initialize_app(firebase_cred)
print(firebase_default_app.name)



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
