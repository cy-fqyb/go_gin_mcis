// app/api/upload_records/route.ts
import { NextResponse } from "next/server";

export async function GET() {
  const records = [
    { id: 1, filename: "report_1.pdf", status: "已完成" },
    { id: 2, filename: "report_2.pdf", status: "处理中" },
  ];
  return NextResponse.json(records);
}
