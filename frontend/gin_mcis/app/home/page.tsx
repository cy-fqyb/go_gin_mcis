"use client";

import { use, useEffect, useState } from "react";
import ReactECharts from "echarts-for-react";

// ---- 模拟接口数据 ----
const fetchPregnantList = () =>
  Promise.resolve([
    { name: "李芳", age: 28, weeks: 12, createTime: "09:12" },
    { name: "王梅", age: 31, weeks: 20, createTime: "09:45" },
    { name: "周倩", age: 26, weeks: 16, createTime: "10:03" },
    { name: "赵丽", age: 29, weeks: 22, createTime: "10:15" },
    { name: "陈娜", age: 34, weeks: 30, createTime: "10:42" },
  ]);

const fetchSummaryData = () =>
  Promise.resolve({
    total: 58,
    today: 12,
    risk: 3,
  });

// ---- 大屏布局 ----
export default function HomPage() {
  const [list, setList] = useState([]as any[]);
  const [summary, setSummary] = useState({ total: 0, today: 0, risk: 0 });

  useEffect(() => {
    fetchPregnantList().then(setList as any);
    fetchSummaryData().then(setSummary);
  }, []);

  const chartOption = {
    textStyle: { color: "#E6F7FF" },
    xAxis: {
      type: "category",
      data: ["一", "二", "三", "四", "五", "六", "日"],
      axisLine: { lineStyle: { color: "#4FC3F7" } },
    },
    yAxis: {
      type: "value",
      axisLine: { lineStyle: { color: "#4FC3F7" } },
      splitLine: { lineStyle: { color: "rgba(255,255,255,0.1)" } },
    },
    series: [
      {
        data: [12, 16, 11, 18, 19, 22, 15],
        type: "line",
        smooth: true,
        lineStyle: { color: "#4FC3F7", width: 3 },
        areaStyle: {
          color: "rgba(79,195,247,0.2)",
        },
      },
    ],
  };

  return (
    <div
      style={{
        height: "100%",
        background: "linear-gradient(180deg, #0A233E 0%, #05101F 100%)",
        color: "#E6F7FF",
        padding: 20,
        boxSizing: "border-box",
      }}
    >
      <header
        style={{
          textAlign: "center",
          fontSize: 32,
          fontWeight: "bold",
          letterSpacing: 4,
          padding: "10px 0",
          color: "#4FC3F7",
          textShadow: "0 0 12px rgba(79,195,247,0.6)",
        }}
      >
        妇产科今日数据大屏
      </header>

      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 2fr 1fr",
          gap: 20,
          height: "calc(100% - 80px)",
        }}
      >
        {/* LEFT MODULE */}
        <section
          style={{
            background: "rgba(15, 36, 64, 0.6)",
            borderRadius: 12,
            border: "1px solid rgba(79,195,247,0.4)",
            padding: 20,
            overflow: "hidden",
            position: "relative",
          }}
        >
          <h3 style={{ fontSize: 20, marginBottom: 10,textAlign: "center" }}>今日建档孕妇名单</h3>

          <div
            style={{
              whiteSpace: "nowrap",
              overflow: "hidden",
              width: "100%",
            }}
          >
            <div
              style={{
                display: "inline-block",
                paddingLeft: "100%",
                animation: "scrollLeft 18s linear infinite",
              }}
            >
              {list.map((item, i) => (
                <div
                  key={i}
                  style={{
                    fontSize: 16,
                    marginRight: 40,
                    display: "inline-block",
                  }}
                >
                  {item.name} · {item.age}岁 · {item.weeks}周 · {item.createTime}
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* CENTER MODULE */}
        <section
          style={{
            background: "rgba(15, 36, 64, 0.6)",
            borderRadius: 12,
            border: "1px solid rgba(79,195,247,0.4)",
            padding: 20,
            overflow: "hidden",
          }}
        >
          <h3 style={{ fontSize: 20, marginBottom: 10 }}>本周建档趋势</h3>
          <ReactECharts option={chartOption} style={{ height: "90%" }} />
        </section>

        {/* RIGHT MODULE */}
        <section
          style={{
            background: "rgba(15, 36, 64, 0.6)",
            borderRadius: 12,
            border: "1px solid rgba(79,195,247,0.4)",
            padding: 20,
            overflow: "hidden",
          }}
        >
          <h3 style={{ fontSize: 20, marginBottom: 10 }}>今日总览</h3>

          <div style={{ fontSize: 22, margin: "20px 0" }}>
            总人数：
            <span style={{ color: "#4FC3F7", fontSize: 28 }}>{summary.total}</span>
          </div>

          <div style={{ fontSize: 22, margin: "20px 0" }}>
            今日建档：
            <span style={{ color: "#4FC3F7", fontSize: 28 }}>{summary.today}</span>
          </div>

          <div style={{ fontSize: 22, margin: "20px 0" }}>
            高风险人数：
            <span style={{ color: "#F06292", fontSize: 28 }}>{summary.risk}</span>
          </div>
        </section>
      </div>

      {/* CSS ANIMATION */}
      <style>{`
        @keyframes scrollLeft {
          0% { transform: translateX(0); }
          100% { transform: translateX(-100%); }
        }
      `}</style>
    </div>
  );
}
