"use client"

import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { Calendar, Users } from "lucide-react"
import { useState } from "react"
import { BookingDialog } from "./booking-dialog"

const courts = [
  {
    id: 1,
    name: "Lapangan A - Premium",
    description: "Lapangan premium dengan rumput sintetis terbaik dan lampu LED",
    price: "Rp 150.000",
    capacity: "10-12 pemain",
    features: ["Rumput Sintetis", "Lampu LED", "AC Ruang Ganti"],
    image: "/modern-indoor-futsal-court-green.jpg",
    available: true,
  },
  {
    id: 2,
    name: "Lapangan B - Standar",
    description: "Lapangan standar dengan fasilitas lengkap dan nyaman",
    price: "Rp 120.000",
    capacity: "10-12 pemain",
    features: ["Rumput Sintetis", "Lampu Standar", "Ruang Ganti"],
    image: "/futsal-court-indoor-lighting.jpg",
    available: true,
  },
  {
    id: 3,
    name: "Lapangan C - Ekonomis",
    description: "Lapangan ekonomis dengan kualitas terjamin",
    price: "Rp 100.000",
    capacity: "10-12 pemain",
    features: ["Rumput Sintetis", "Lampu Standar"],
    image: "/outdoor-futsal-field-night.jpg",
    available: true,
  },
  {
    id: 4,
    name: "Lapangan D - VIP",
    description: "Lapangan VIP dengan tribun penonton dan fasilitas eksklusif",
    price: "Rp 200.000",
    capacity: "10-12 pemain",
    features: ["Rumput Import", "Lampu Premium", "Tribun", "Kafe"],
    image: "/vip-futsal-arena-tribune.jpg",
    available: false,
  },
  {
    id: 5,
    name: "Lapangan E - Mini",
    description: "Lapangan mini ideal untuk latihan atau tim kecil",
    price: "Rp 80.000",
    capacity: "6-8 pemain",
    features: ["Rumput Sintetis", "Lampu LED"],
    image: "/mini-futsal-court-practice.jpg",
    available: true,
  },
]

export function CourtList() {
  const [selectedCourt, setSelectedCourt] = useState<(typeof courts)[0] | null>(null)
  const [showBookingDialog, setShowBookingDialog] = useState(false)

  const handleBooking = (court: (typeof courts)[0]) => {
    setSelectedCourt(court)
    setShowBookingDialog(true)
  }

  return (
    <>
      <section id="lapangan" className="scroll-mt-16 sm:scroll-mt-20">
        <div className="mb-6 sm:mb-8 lg:mb-10">
          <h2 className="text-2xl sm:text-3xl lg:text-4xl font-bold mb-2 sm:mb-3">Pilih Lapangan</h2>
          <p className="text-muted-foreground text-sm sm:text-base lg:text-lg">5 lapangan dengan berbagai pilihan dan fasilitas</p>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 gap-4 sm:gap-5 lg:gap-6">
          {courts.map((court) => (
            <Card key={court.id} className="overflow-hidden hover:shadow-lg transition-all duration-300 hover:scale-[1.02]">
              <div className="relative h-40 sm:h-44 lg:h-48 overflow-hidden bg-muted">
                <img 
                  src={court.image || "/placeholder.svg"} 
                  alt={court.name} 
                  className="w-full h-full object-cover transition-transform duration-300 hover:scale-105" 
                />
                {!court.available && (
                  <div className="absolute inset-0 bg-background/80 flex items-center justify-center">
                    <Badge variant="secondary" className="text-xs sm:text-sm">
                      Sedang Digunakan
                    </Badge>
                  </div>
                )}
              </div>

              <CardHeader className="p-3 sm:p-4 lg:p-6">
                <div className="flex flex-col sm:flex-row sm:items-start justify-between gap-2">
                  <CardTitle className="text-lg sm:text-xl leading-tight">{court.name}</CardTitle>
                  <Badge variant="outline" className="shrink-0 w-fit">
                    {court.price}
                  </Badge>
                </div>
                <CardDescription className="text-xs sm:text-sm leading-relaxed">{court.description}</CardDescription>
              </CardHeader>

              <CardContent className="space-y-3 p-3 sm:p-4 lg:p-6 pt-0">
                <div className="flex items-center gap-2 text-xs sm:text-sm text-muted-foreground">
                  <Users className="w-3 h-3 sm:w-4 sm:h-4" />
                  <span>{court.capacity}</span>
                </div>

                <div className="flex flex-wrap gap-1 sm:gap-2">
                  {court.features.map((feature, index) => (
                    <Badge key={index} variant="secondary" className="text-xs px-2 py-1">
                      {feature}
                    </Badge>
                  ))}
                </div>
              </CardContent>

              <CardFooter className="p-3 sm:p-4 lg:p-6 pt-0">
                <Button 
                  className="w-full text-xs sm:text-sm py-2 sm:py-2.5" 
                  disabled={!court.available} 
                  onClick={() => handleBooking(court)}
                >
                  <Calendar className="w-3 h-3 sm:w-4 sm:h-4 mr-1 sm:mr-2" />
                  <span className="hidden sm:inline">Booking Sekarang</span>
                  <span className="sm:hidden">Book</span>
                </Button>
              </CardFooter>
            </Card>
          ))}
        </div>
      </section>

      {selectedCourt && (
        <BookingDialog open={showBookingDialog} onOpenChange={setShowBookingDialog} court={selectedCourt} />
      )}
    </>
  )
}
