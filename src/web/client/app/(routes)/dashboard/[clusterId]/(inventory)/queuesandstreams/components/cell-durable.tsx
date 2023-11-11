import { RabbitMqQueue } from "@/models/queues";

interface CellDurableProps {
  Queue: RabbitMqQueue;
}
function CellDurable({ Queue }: CellDurableProps) {
  if (Queue.Durable) return <p>Durable</p>;
  return <p>Non Durable</p>;
}

export default CellDurable;
