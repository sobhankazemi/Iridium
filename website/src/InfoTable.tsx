import React, { useState, useEffect } from "react";
import axios from "axios";
import {
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  getKeyValue,
} from "@nextui-org/react";

interface Data {
  os: string;
  kernelName: string;
  hostName: string;
  kernelRelease: string;
  kernelVersion: string;
  machine: string;
  processor: string;
  hwPlatform: string;
  usedSpace: string;
  dateTime: string;
}

const InfoTable: React.FC = () => {
  const [data, setData] = useState<Data[]>([]);

  const fetchData = async () => {
    const result = await axios.get(
      // "https://jsonplaceholder.typicode.com/users"
      "attacker:8080/info"
    );
    setData(result.data);
  };

  useEffect(() => {
    fetchData();
  }, []);

  useEffect(() => {
    const interval = setInterval(() => {
      fetchData();
    }, 5000);

    return () => clearInterval(interval);
  }, []);
  const columns = [
    { Header: "OS", accessor: "os" },
    { Header: "KernelName", accessor: "kernelName" },
    { Header: "HostName", accessor: "hostName" },
    { Header: "KernelRelease", accessor: "kernelRelease" },
    { Header: "KernelVersion", accessor: "kernelVersion" },
    { Header: "Machine", accessor: "machine" },
    { Header: "Processor", accessor: "processor" },
    { Header: "HWPlatform", accessor: "hwPlatform" },
    { Header: "UsedSpace", accessor: "usedSpace" },
    { Header: "DateTime", accessor: "dateTime" },
  ];
  return (
    <Table aria-label="Example table with dynamic content">
      <TableHeader>
        {columns.map((column) => (
          <TableColumn key={column.accessor}>{column.Header}</TableColumn>
        ))}
      </TableHeader>
      <TableBody emptyContent={"No rows to display."}>
        {data.map((row) => (
          <TableRow key={row.dateTime}>
            {(columnKey) => (
              <TableCell>{getKeyValue(row, columnKey)}</TableCell>
            )}
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

export default InfoTable;
