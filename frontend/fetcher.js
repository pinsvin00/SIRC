export let guid = "";
export const ip = "your ip here";
export function set_credentials(credentials) {
  guid = credentials;
}

export async function fetch_messages() {
  console.log("listening!");
  let data = {};
  data.guid = guid;
  let response = await fetch(ip + "/subscribe", {
    method: "POST",
    body: JSON.stringify(data),
  });
  if (response.status === 502) {
    await fetch();
  }

  const messages = await response.json();
  messages.forEach((message) => {
    addMessage(message, true);
  });

  await fetch_messages();
}

export async function send_message(text) {
  let data = {};
  data.guid = guid;
  data.text = text;

  const response = await fetch(ip + "/send_message", {
    method: "POST",
    body: JSON.stringify(data),
  });

  const json = await response.json();
  if (json.success === "false") {
    addMessage({ text: "Nie jestes upowazniony do komunikacji w tym kanale!", color: "red", sender: "Server" });
  }

  if (document.getElementById("message").value.startsWith("/quit")) {
    document.cookie.split(";").forEach(function (c) {
      document.cookie = c.replace(/^ +/, "").replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
    });
    window.location.href = ip + "/index.html";
  }

  document.getElementById("message").value = "";
}

function addMessage(message, isUsers) {
  console.log("new message incoming!", message);
  const new_message = document.createElement("div");
  const main_container = document.getElementById("messages");
  new_message.innerHTML = `      
  <div class="container">
      <img
        src="https://upload.wikimedia.org/wikipedia/commons/7/7c/Profile_avatar_placeholder_large.png"
        alt="Avatar"
        class="left"
        style="width: 100%"
      />
      <p class="emoticonize">${message.text}</p>
      <div class="time-right">
      <span class="mr-3" style="color:${message.color}">${message.sender}</span>
      <span class="time-right" style="margin-left: 1rem">${message.time}</span>
    </div>
  </div>`;
  main_container.appendChild(new_message);
}
