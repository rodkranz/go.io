# Socket.io Go

#### Start Server
    
    $ go run main.go
    
> you can access by url: `http://localhost:5000/`
 
 
---
#### Activate WebHook

    $ http://localhost:5000/webhook?act=Chat&txt=Lorem
    
 
> The `act` (activation) and `txt` (text) are mandatory fields
 
 

---
#### Frontend Javascript
    
    (function () {
        'use strict';

        var socket = io();
        var act = 'Chat'; // activation
        // Receive data from socket and call `receiveData`
        socket.on(act, receiveData);

        /**
         * ReceiveData
         * @param data
         */
        function receiveData(data) {
            var msg = message(data);
            $('#messages').append(msg);
            animation(msg);
        }

        /**
         * Create element li empty with display none
         * @param data
         * @returns {*|jQuery}
         */
        function message(data) {
            return $('<li />')
                    .css({display: "none"})
                    .text(data);
        }

        /**
         * Start animation
         * @param elm
         * @param callback
         */
        function animation(elm, callback) {
            var self = $(elm);
            self.fadeIn(700, function () {
                setTimeout(function() {
                    self.fadeOut(700, callback)
                }, 5000);
            });
        }
    })();
---
    