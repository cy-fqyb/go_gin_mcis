import { log } from "console";

// app/upload_record/page.tsx
export default async function UploadRecord() {
const res = await fetch("http://localhost:3000/api/upload_records", { method: "GET", cache: "no-store" });



  log(res);
  const records = await res.json();

  return (
    <div className="p-8">
      <h1 className="text-2xl font-bold mb-4">上传记录列表</h1>
      <ul className="space-y-2">
        {records.map((r: any) => (
          <li key={r.id} className="border p-3 rounded">
            {r.filename} - {r.status}
          </li>
        ))}
      </ul>
    </div>
  );
}
