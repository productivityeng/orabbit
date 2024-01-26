import { RabbitMqExchange } from "@/models/exchange";

interface CellDurableProps {
  Data: RabbitMqExchange;
}
function CellDurable({ Data }: CellDurableProps) {
  if (Data.Durable) return <p>Durable</p>;
  return <p>Non Durable</p>;
}

export default CellDurable;
