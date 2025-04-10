import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { ArrowLeft, Download } from "lucide-react";
import Link from "next/link";

export default function InvoiceDetailsPage({
  params,
}: {
  params: { id: string };
}) {
  const { id } = params;
  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Button
          variant="ghost"
          className="text-gray-400 hover:text-white"
          asChild
        >
          <Link href="/invoices">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Voltar
          </Link>
        </Button>

        <div className="flex-1">
          <div className="flex items-center gap-3">
            <h1 className="text-2xl font-bold text-white">Fatura {id}</h1>
            <Badge
              variant="outline"
              className="bg-green-500/20 text-green-500 hover:bg-green-500/20"
            >
              Aprovado
            </Badge>
          </div>
          <p className="text-gray-400">
            Criada em {new Date().toLocaleDateString()}
          </p>
        </div>

        <Button variant="outline" className="bg-[#2a3749] border-gray-700">
          <Download className="h-4 w-4 mr-2" />
          Download PDF
        </Button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {/* Informações da Fatura */}
        <Card className="bg-[#1e293b] border-gray-800 p-6">
          <h2 className="text-xl font-semibold text-white mb-4">
            Informações da Fatura
          </h2>

          <div className="space-y-4">
            <div className="flex justify-between border-b border-gray-800 pb-2">
              <span className="text-gray-400">ID da Fatura</span>
              <span className="text-white font-medium">{id}</span>
            </div>

            <div className="flex justify-between border-b border-gray-800 pb-2">
              <span className="text-gray-400">Valor</span>
              <span className="text-white font-medium">
                R$ 0,00
              </span>
            </div>

            <div className="flex justify-between border-b border-gray-800 pb-2">
              <span className="text-gray-400">Data de Criação</span>
              <span className="text-white font-medium">
                {new Date().toLocaleDateString()}
              </span>
            </div>

            <div className="flex justify-between pb-2">
              <span className="text-gray-400">Descrição</span>
              <span className="text-white font-medium">
                Fatura de exemplo
              </span>
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
}
