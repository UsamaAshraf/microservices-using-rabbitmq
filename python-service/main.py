from flask import Flask
from flask import request

app = Flask(__name__)

@app.route('/users/<int:user_id>', methods=['POST'])
def update(user_id):
    new_name = request.form['full_name']
    #TODO: Emit the event...
    return make_response({full_name: new_name}, 201)
