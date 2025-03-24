$('#create-user-form').on('submit', createUser)

function createUser(event) {
  event.preventDefault();

  if ($('#password').val() != $('#confirm-password').val()) {
    alert("Password do not match.");
    return
  }

  $.ajax({
    url: '/users',
    method: 'POST',
    data: {
      name: $('#name').val(),
      nick: $('#nick').val(),
      email: $('#email').val(),
      password: $('#password').val()
    }
  }).done(function() {
    alert('User created successfully.');
  }).fail(function(erro) {
    console.log(erro);
    alert('Error creating user.');
  });
}