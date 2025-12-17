"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import { Calendar } from "@/components/ui/calendar"
import { Badge } from "@/components/ui/badge"
import { CalendarIcon, Clock } from "lucide-react"

interface BookingDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  court: {
    id: number
    name: string
    price: string
    capacity: string
  }
}

const timeSlots = [
  "08:00 - 09:00",
  "09:00 - 10:00",
  "10:00 - 11:00",
  "11:00 - 12:00",
  "12:00 - 13:00",
  "13:00 - 14:00",
  "14:00 - 15:00",
  "15:00 - 16:00",
  "16:00 - 17:00",
  "17:00 - 18:00",
  "18:00 - 19:00",
  "19:00 - 20:00",
  "20:00 - 21:00",
  "21:00 - 22:00",
]

export function BookingDialog({ open, onOpenChange, court }: BookingDialogProps) {
  const [date, setDate] = useState<Date | undefined>(new Date())
  const [selectedTime, setSelectedTime] = useState<string>("")
  const [bookedSlots] = useState<string[]>(["10:00 - 11:00", "15:00 - 16:00", "19:00 - 20:00"])
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  const handleBooking = () => {
    if (!date || !selectedTime) {
      alert("Pilih tanggal dan waktu booking!")
      return
    }

    // TODO: Implement booking logic
    console.log("Booking:", { court, date, time: selectedTime })
    alert(
      `Booking berhasil!\nLapangan: ${court.name}\nTanggal: ${date.toLocaleDateString("id-ID")}\nWaktu: ${selectedTime}`,
    )
    onOpenChange(false)
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="w-[95vw] max-w-xs sm:max-w-lg lg:max-w-3xl xl:max-w-4xl max-h-[90vh] overflow-y-auto mx-auto">
        <DialogHeader className="space-y-2 sm:space-y-3">
          <DialogTitle className="text-lg sm:text-xl lg:text-2xl leading-tight">{court.name}</DialogTitle>
          <DialogDescription className="text-sm sm:text-base">
            Pilih tanggal dan waktu untuk booking lapangan
          </DialogDescription>
        </DialogHeader>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6 mt-3 sm:mt-4">
          <div className="space-y-3 sm:space-y-4 order-2 lg:order-1">
            <div>
              <Label className="text-sm sm:text-base mb-2 sm:mb-3 flex items-center gap-2">
                <CalendarIcon className="w-3 h-3 sm:w-4 sm:h-4" />
                Pilih Tanggal
              </Label>
              <div className="border rounded-lg p-2 sm:p-3 bg-card">
                {mounted ? (
                  <Calendar
                    mode="single"
                    selected={date}
                    onSelect={setDate}
                    disabled={(date) => date < new Date()}
                    className="rounded-md w-full"
                  />
                ) : (
                  <div className="h-64 bg-muted rounded-md animate-pulse" />
                )}
              </div>
            </div>

            <div className="p-3 sm:p-4 bg-muted rounded-lg space-y-2">
              <h3 className="font-semibold text-sm sm:text-base">Detail Lapangan</h3>
              <div className="space-y-1 text-xs sm:text-sm">
                <p className="text-muted-foreground">
                  Harga: <span className="text-foreground font-medium">{court.price}</span>/jam
                </p>
                <p className="text-muted-foreground">
                  Kapasitas: <span className="text-foreground font-medium">{court.capacity}</span>
                </p>
              </div>
            </div>
          </div>

          <div className="space-y-3 sm:space-y-4 order-1 lg:order-2">
            <Label className="text-sm sm:text-base flex items-center gap-2">
              <Clock className="w-3 h-3 sm:w-4 sm:h-4" />
              Pilih Waktu
            </Label>
            <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-2 xl:grid-cols-3 gap-2 max-h-64 sm:max-h-80 lg:max-h-96 overflow-y-auto pr-2">
              {timeSlots.map((slot) => {
                const isBooked = bookedSlots.includes(slot)
                const isSelected = selectedTime === slot

                return (
                  <button
                    key={slot}
                    onClick={() => !isBooked && setSelectedTime(slot)}
                    disabled={isBooked}
                    className={`
                      p-2 sm:p-3 rounded-lg border text-xs sm:text-sm font-medium transition-all hover:scale-105 touch-target
                      ${
                        isSelected
                          ? "bg-primary text-primary-foreground border-primary"
                          : isBooked
                            ? "bg-muted text-muted-foreground border-border cursor-not-allowed opacity-50"
                            : "bg-card hover:bg-accent border-border hover:border-primary"
                      }
                    `}
                  >
                    <span className="block">{slot}</span>
                    {isBooked && (
                      <Badge variant="secondary" className="mt-1 text-xs">
                        Penuh
                      </Badge>
                    )}
                  </button>
                )
              })}
            </div>
          </div>
        </div>

        <div className="flex flex-col sm:flex-row items-center justify-between gap-3 sm:gap-4 pt-3 sm:pt-4 border-t">
          <div className="w-full sm:w-auto">
            {date && selectedTime && (
              <p className="text-xs sm:text-sm text-muted-foreground text-center sm:text-left">
                Booking:{" "}
                <span className="text-foreground font-medium block sm:inline">
                  {date.toLocaleDateString("id-ID")} - {selectedTime}
                </span>
              </p>
            )}
          </div>
          <Button 
            size="lg" 
            onClick={handleBooking} 
            disabled={!date || !selectedTime}
            className="w-full sm:w-auto h-10 sm:h-11 text-sm sm:text-base"
          >
            Konfirmasi Booking
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}
