import { Button } from "@/components/ui/button";
import { ShoppingCart } from "lucide-react"

interface BookCardProps {
  imgUrl: string,
  title: string,
  author: string,
  price: number,
}

function formattedPrice(x: number) {
  return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

export function BookCard({ imgUrl, title, author, price }: BookCardProps) {
  return (
    <div className="w-52 max-w-md m-5">
      <div className="bg-white rounded-lg shadow-lg overflow-hidden transition-all duration-300 hover:shadow-xl dark:bg-gray-950">
        <img
          src={imgUrl}
          alt={title}
          width={600}
          height={400}
          className="w-50 max-w-60 object-cover"
          style={{ aspectRatio: "400/450", objectFit: "cover" }}
        />
        <div className="p-4 space-y-2">
          <h3 className="text-base font-semibold truncate">{title}</h3>
          <p className="text-sm text-gray-800 dark:text-gray-400">{author}</p>
          <div className="flex items-center justify-between">
            <span className="text-base font-bold">Rp{formattedPrice(price)}</span>
            <Button size="icon">
              <ShoppingCart />
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}