"use client"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Clock, CalendarIcon } from "lucide-react"
import { useState, useEffect } from "react"
import { Calendar } from "@/components/ui/calendar"

const courtSchedule = [
  { court: "Lapangan A", time: "08:00 - 09:00", status: "available" },
  { court: "Lapangan A", time: "09:00 - 10:00", status: "booked" },
  { court: "Lapangan A", time: "10:00 - 11:00", status: "booked" },
  { court: "Lapangan A", time: "11:00 - 12:00", status: "available" },

  { court: "Lapangan B", time: "08:00 - 09:00", status: "available" },
  { court: "Lapangan B", time: "09:00 - 10:00", status: "available" },
  { court: "Lapangan B", time: "10:00 - 11:00", status: "booked" },
  { court: "Lapangan B", time: "11:00 - 12:00", status: "available" },

  { court: "Lapangan C", time: "08:00 - 09:00", status: "booked" },
  { court: "Lapangan C", time: "09:00 - 10:00", status: "available" },
  { court: "Lapangan C", time: "10:00 - 11:00", status: "available" },
  { court: "Lapangan C", time: "11:00 - 12:00", status: "booked" },

  { court: "Lapangan D", time: "08:00 - 09:00", status: "available" },
  { court: "Lapangan D", time: "09:00 - 10:00", status: "booked" },
  { court: "Lapangan D", time: "10:00 - 11:00", status: "available" },
  { court: "Lapangan D", time: "11:00 - 12:00", status: "available" },

  { court: "Lapangan E", time: "08:00 - 09:00", status: "available" },
  { court: "Lapangan E", time: "09:00 - 10:00", status: "available" },
  { court: "Lapangan E", time: "10:00 - 11:00", status: "booked" },
  { court: "Lapangan E", time: "11:00 - 12:00", status: "available" },
]

const courts = ["Lapangan A", "Lapangan B", "Lapangan C", "Lapangan D", "Lapangan E"]
const timeSlots = ["08:00 - 09:00", "09:00 - 10:00", "10:00 - 11:00", "11:00 - 12:00"]

export function AvailabilityCalendar() {
  const [selectedDate, setSelectedDate] = useState<Date | undefined>(new Date())
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  const getStatus = (court: string, time: string) => {
    const schedule = courtSchedule.find((s) => s.court === court && s.time === time)
    return schedule?.status || "available"
  }

  return (
    <section id="ketersediaan" className="scroll-mt-16 sm:scroll-mt-20">
      <Card>
        <CardHeader className="p-4 sm:p-6">
          <CardTitle className="text-xl sm:text-2xl lg:text-3xl flex flex-col sm:flex-row items-start sm:items-center gap-2 sm:gap-3">
            <CalendarIcon className="w-6 h-6 sm:w-7 sm:h-7 lg:w-8 lg:h-8 text-primary" />
            <span className="leading-tight">Lihat Ketersediaan Lapangan</span>
          </CardTitle>
          <CardDescription className="text-sm sm:text-base lg:text-lg">
            Cek jadwal ketersediaan semua lapangan secara real-time
          </CardDescription>
        </CardHeader>

        <CardContent className="space-y-4 sm:space-y-6 p-4 sm:p-6">
          <div className="grid grid-cols-1 lg:grid-cols-[280px_1fr] xl:grid-cols-[320px_1fr] gap-4 sm:gap-6">
            <div className="border rounded-lg p-3 sm:p-4 bg-card order-2 lg:order-1">
              {mounted ? (
                <Calendar
                  mode="single"
                  selected={selectedDate}
                  onSelect={setSelectedDate}
                  disabled={(date) => date < new Date()}
                  className="rounded-md w-full"
                />
              ) : (
                <div className="h-64 bg-muted rounded-md animate-pulse" />
              )}
            </div>

            <div className="space-y-3 sm:space-y-4 order-1 lg:order-2">
              <div className="flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-4 pb-2 border-b">
                <Badge variant="outline" className="bg-primary/10 border-primary w-fit">
                  <div className="w-2 h-2 rounded-full bg-primary mr-2" />
                  Tersedia
                </Badge>
                <Badge variant="outline" className="bg-muted w-fit">
                  <div className="w-2 h-2 rounded-full bg-muted-foreground mr-2" />
                  Terisi
                </Badge>
              </div>

              <div className="overflow-x-auto -mx-2 sm:mx-0">
                <div className="min-w-[600px] sm:min-w-full">
                  <table className="w-full border-collapse">
                    <thead>
                      <tr className="border-b">
                        <th className="text-left p-2 sm:p-3 font-semibold sticky left-0 bg-card z-10 min-w-[100px] sm:min-w-[120px]">
                          <Clock className="w-3 h-3 sm:w-4 sm:h-4 inline mr-1 sm:mr-2" />
                          <span className="text-xs sm:text-sm">Waktu</span>
                        </th>
                        {courts.map((court) => (
                          <th key={court} className="text-center p-2 sm:p-3 font-semibold min-w-[90px] sm:min-w-[100px] lg:min-w-[120px]">
                            <span className="text-xs sm:text-sm">{court}</span>
                          </th>
                        ))}
                      </tr>
                    </thead>
                    <tbody>
                      {timeSlots.map((time) => (
                        <tr key={time} className="border-b hover:bg-accent/50 transition-colors">
                          <td className="p-2 sm:p-3 font-medium sticky left-0 bg-card text-xs sm:text-sm">
                            {time}
                          </td>
                          {courts.map((court) => {
                            const status = getStatus(court, time)
                            return (
                              <td key={`${court}-${time}`} className="p-1.5 sm:p-2 lg:p-3 text-center">
                                <div
                                  className={`
                                    inline-flex items-center justify-center px-2 sm:px-3 py-1 sm:py-1.5 rounded-md text-xs sm:text-sm font-medium transition-all hover:scale-105 cursor-pointer
                                    ${
                                      status === "available"
                                        ? "bg-primary/10 text-primary border border-primary/20 hover:bg-primary/20"
                                        : "bg-muted text-muted-foreground border border-border hover:bg-muted/80"
                                    }
                                  `}
                                >
                                  <span className="hidden sm:inline">
                                    {status === "available" ? "Tersedia" : "Terisi"}
                                  </span>
                                  <span className="sm:hidden">
                                    {status === "available" ? "✓" : "×"}
                                  </span>
                                </div>
                              </td>
                            )
                          })}
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>

          <div className="p-3 sm:p-4 bg-muted rounded-lg">
            <p className="text-xs sm:text-sm text-muted-foreground text-center leading-relaxed">
              Menampilkan ketersediaan untuk{" "}
              <span className="text-foreground font-medium block sm:inline mt-1 sm:mt-0">
                {selectedDate?.toLocaleDateString("id-ID", {
                  weekday: "long",
                  year: "numeric",
                  month: "long",
                  day: "numeric",
                })}
              </span>
            </p>
          </div>
        </CardContent>
      </Card>
    </section>
  )
}
