/*
 * @Author: your name
 * @Date: 2020-12-11 11:55:49
 * @LastEditTime: 2020-12-11 11:57:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/static/js/websocket.js
 */
var socket;

$(document).ready(function () {
    // Create a socket
    socket = new WebSocket('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text());
    // Message received on the socket
    socket.onmessage = function (event) {
        var data = JSON.parse(event.data);
        var li = document.createElement('li');

        console.log(data);

        switch (data.Type) {
        case 0: // JOIN
            if (data.User == $('#uname').text()) {
                li.innerText = '你加入了聊天室。';
            } else {
                li.innerText = data.User + ' 加入了聊天室。';
            }
            break;
        case 1: // LEAVE
            li.innerText = data.User + ' 离开了聊天室。';
            break;
        case 2: // MESSAGE
            var username = document.createElement('strong');
            var content = document.createElement('span');

            username.innerText = data.User;
            content.innerText = data.Content;

            li.appendChild(username);
            li.appendChild(document.createTextNode(': '));
            li.appendChild(content);

            break;
        }

        $('#chatbox li').first().before(li);
    };

    // Send messages.
    var postConecnt = function () {
        var uname = $('#uname').text();
        var content = $('#sendbox').val();
        socket.send(content);
        $('#sendbox').val('');
    }

    $('#sendbtn').click(function () {
        postConecnt();
    });
});
