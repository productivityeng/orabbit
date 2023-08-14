import React from "react";

interface SimpleHeadingProps {
  title: string;
  description: string;
}
const SimpleHeading = ({ title, description }: SimpleHeadingProps) => {
  return (
    <div>
      <h2 className="text-3xl font-bold tracking-tight">{title}</h2>
      <p className="text-muted-foreground text-sm">{description}</p>
    </div>
  );
};

export default SimpleHeading;
