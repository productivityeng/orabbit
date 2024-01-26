import { RabbitMqQueue } from "@/models/queues";
import { RabbitMqUser } from "@/models/users";
import React from "react";

interface CellArgumentsProps {
  User: RabbitMqUser;
}
function CellArguments({ User }: CellArgumentsProps) {
  return <b>{JSON.stringify(User)}</b>;
}

export default CellArguments;
