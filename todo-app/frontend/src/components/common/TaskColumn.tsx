// components/common/TaskColumn.tsx
import React from "react";
import { FiTrash2 } from "react-icons/fi";

interface Task {
  id: number;
  title: string;
  category: string;
  assignee: string;
  deadline: string;
  description: string;
  done: boolean;
}

interface TaskColumnProps {
  columnId: string;
  column: {
    name: string;
    items: Task[];
  };
  onTaskClick: (columnId: string, taskId: number) => void; // Geändert zu taskId
  onDeleteClick: (columnId: string, taskId: number) => void; // Geändert zu taskId
  onAddTask: (columnId: string) => void;
}

const TaskColumn: React.FC<TaskColumnProps> = ({
  columnId,
  column,
  onTaskClick,
  onDeleteClick,
  onAddTask,
}) => (
  <div className="bg-neutral-900 p-4 rounded shadow-md">
    <h2 className="font-semibold mb-2">{column.name}</h2>
    <ul>
      {column.items.map((item) => (
        <li
          key={item.id}
          className="bg-neutral-950 text-neutral-400 p-2 mb-2 rounded shadow flex justify-between items-center hover:bg-neutral-900 cursor-pointer"
        >
          <span onClick={() => onTaskClick(columnId, item.id)}>
            {item.title}
          </span>
          <button
            onClick={() => onDeleteClick(columnId, item.id)}
            className="text-red-500 hover:text-red-600 transition-colors"
          >
            <FiTrash2 />
          </button>
        </li>
      ))}
    </ul>
    <button
      className="mt-2 bg-info text-white px-4 py-2 rounded"
      onClick={() => onAddTask(columnId)}
    >
      + Neue Aufgabe
    </button>
  </div>
);

export default TaskColumn;
