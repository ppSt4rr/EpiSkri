let ws = null;
let username = null;
let userId = null;

function submitUsername() {
    const inputUsername = document.getElementById("username").value;
    if (inputUsername) {
        username = inputUsername;

        ws = new WebSocket('ws://localhost:8080/websocket');

        ws.onopen = function() {
            const message = {
                type: 'setUsername',
                username: username
            };
            ws.send(JSON.stringify(message));
        };

        ws.onmessage = function(event) {
            try {
                const response = JSON.parse(event.data);

                switch (response.type) {
                    case "userId":
                        userId = response.userId;
                        console.log("Votre ID utilisateur est : " + userId);
                        break;
                    case "message":
                        addMessageToContainer(response.username, response.message);
                        break;
                    case "drawing":
                        drawOnCanvas(response.x, response.y, response.color, response.thickness);
                        break;
                    case "userList":
                        if (Array.isArray(response.userList)) {
                            updateUserList(response.userList);
                        } else {
                            console.error("Le format de la liste des utilisateurs est incorrect :", response.userList);
                        }
                        break;
                    default:
                        console.error("Type de message non reconnu :", response.type);
                }
            } catch (error) {
                console.error("Erreur lors du traitement du message reçu :", error);
            }
        };

        ws.onerror = function(error) {
            console.error("Erreur WebSocket :", error);
        };

        document.getElementById("usernameContainer").style.display = "none";
        document.getElementById("chatContainer").style.display = "block";
    } else {
        alert("Veuillez entrer un nom d'utilisateur !");
    }
}

function updateUserList(userList) {
    const userContainer = document.getElementById("userList");
    userContainer.innerHTML = ""; 

    if (userList.length === 0) {
        const noUsers = document.createElement("p");
        noUsers.innerText = "Aucun utilisateur connecté.";
        userContainer.appendChild(noUsers);
        return;
    }

    userList.forEach(user => {
        const userElement = document.createElement("p");
        userElement.innerText = `${user.username} (ID: ${user.userId})`;
        userContainer.appendChild(userElement);
    });
}

function sendMessage() {
    const message = document.getElementById("message").value;

    if (username && message && ws) {
        const data = { type: "message", username: username, message: message };
        ws.send(JSON.stringify(data));
    }
}

function addMessageToContainer(username, message) {
    const newMessage = document.createElement("p");
    newMessage.innerText = `${username}: ${message}`;
    const responseContainer = document.getElementById("responseContainer");
    responseContainer.appendChild(newMessage);
}

const canvas = document.getElementById("drawingCanvas");
const ctx = canvas.getContext("2d");
let drawing = false;
let color = document.getElementById("color").value;
let thickness = document.getElementById("thickness").value;

document.getElementById("color").addEventListener("input", function() {
    color = this.value;
});

document.getElementById("thickness").addEventListener("input", function() {
    thickness = this.value;
});

canvas.addEventListener("mousedown", function() {
    drawing = true;
    ctx.beginPath();
});

canvas.addEventListener("mouseup", function() {
    drawing = false;
});

canvas.addEventListener("mousemove", function(event) {
    if (!drawing) return;

    const rect = canvas.getBoundingClientRect();
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;

    drawOnCanvas(x, y, color, thickness);

    if (username && ws) {
        const drawingData = {
            type: "drawing",
            username: username,
            x: x,
            y: y,
            color: color,
            thickness: thickness
        };
        ws.send(JSON.stringify(drawingData));
    }
});

function drawOnCanvas(x, y, color, thickness) {
    ctx.strokeStyle = color;
    ctx.lineWidth = thickness;
    ctx.lineTo(x, y);
    ctx.stroke();
}
