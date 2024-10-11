from flask import Flask, request, jsonify
import subprocess

app = Flask(__name__)

@app.route("/stackmachine", methods=['POST'])
def execute_commands():
    data = request.get_json()
    commands = data.get("commands", "")

    try:
        result = subprocess.run(
            ["../stackmachine.exe", commands],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True,
            check=True
        )
        output = result.stdout.strip()
        return jsonify({"result": output}), 200
    except subprocess.CalledProcessError as e:
        return jsonify({"error": e.stderr.strip()}), 400


if __name__ == "__main__":
    app.run(debug=True)