function submitForm() {
  let name = document.getElementById("name").value;
  let email = document.getElementById("email").value;
  let phone = document.getElementById("phone").value;
  let subject = document.getElementById("subject").value;
  let message = document.getElementById("message").value;

  if (name == "") {
    return alert("Name input fields must be not empty");
  } else if (email == "") {
    return alert("Email input fields must be not empty");
  } else if (phone == "") {
    return alert("Phone input fields must be not empty");
  } else if (subject == "") {
    return alert("Subject input fields must be not empty");
  } else if (message == "") {
    return alert("Message input fields must be not empty");
  }

  let emailReciever = "akanime1@gmail.com";

  let a = document.createElement("a");

  a.href = `mailto:${emailReciever}?subject=${subject}&body=Hello I'm ${name}, ${subject}, ${message}`;
  a.target = "_blank";
  a.click();

  let dataObject = {
    name: name,
    email: email,
    phone: phone,
    subject: subject,
    message: message,
  };
}


