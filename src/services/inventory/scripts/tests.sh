curl -i -u radmin:radmin http://localhost:15671/api/exchanges

curl -i -u radmin:radmin http://localhost:15671/api/vhosts

curl -i -u radmin:radmin http://localhost:15671/api/queues/%2f/master-queue/bindings 
[{"source":"","vhost":"/","destination":"master-queue","destination_type":"queue","routing_key":"master-queue","arguments":{},"properties_key":"master-queue"},{"source":"exchange-master","vhost":"/","destination":"master-queue","destination_type":"queue","routing_key":"#.d.#.t","arguments":{},"properties_key":"%23.d.%23.t"}]%                                     



{"arguments":{"alternate-exchange":""},"auto_delete":false,"durable":true,"internal":false,"name":"new-exchange-for-testing","type":"fanout","user_who_performed_action":"radmin","vhost":"/"}


