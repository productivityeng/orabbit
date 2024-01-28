import { RabbitMqExchange } from "@/models/exchange";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import React from "react";

interface DefaultQueueTypeProps {
  Data: RabbitMqVirtualHost;
}
function DefaultQueueType({ Data }: DefaultQueueTypeProps) {
  return <b>{Data.DefaultQueueType}</b>;
}

export default DefaultQueueType;
