import { RabbitMqQueue } from "@/models/queues";
import React from "react";

interface CellArgumentsProps {
  Queue: RabbitMqQueue;
}
function CellArguments({ Queue }: CellArgumentsProps) {
  return <b>{JSON.stringify(Queue.Arguments)}</b>;
}

export default CellArguments;
