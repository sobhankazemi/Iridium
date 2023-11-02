import React from "react";
import InfoTable from "./InfoTable";

interface Props {
  title: string;
  description: string;
}

const App: React.FC<Props> = () => {
  return (
    <div>
      <InfoTable />
    </div>
  );
};

export default App;
