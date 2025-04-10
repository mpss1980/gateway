"use client";

import { Button } from "@/components/ui/button";
// import { Input } from "@/components/ui/input"
// import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { PlusIcon, Eye, Download } from "lucide-react";
import Link from "next/link";
import { StatusBadge } from "@/components/StatusBadge";

export function InvoiceList() {
  return (
    <div className="bg-[#1e293b] rounded-lg p-6 border border-gray-800">
      <div className="flex justify-between items-center mb-4">
        <div>
          <h1 className="text-2xl font-bold text-white mb-1">Faturas</h1>
          <p className="text-gray-400">
            Gerencie suas faturas e acompanhe os pagamentos
          </p>
        </div>
        <Button
          className="bg-indigo-600 hover:bg-indigo-700 text-white"
          asChild
        >
          <Link href="/invoices/create">
            <PlusIcon className="h-4 w-4 mr-2" />
            Nova Fatura
          </Link>
        </Button>
      </div>

      {/* Tabela de Faturas */}
      <div className="overflow-x-auto">
        <table className="w-full">
          <thead>
            <tr className="border-b border-gray-800">
              <th className="text-left py-3 px-4 text-gray-400 font-medium">ID</th>
              <th className="text-left py-3 px-4 text-gray-400 font-medium">Valor</th>
              <th className="text-left py-3 px-4 text-gray-400 font-medium">Status</th>
              <th className="text-left py-3 px-4 text-gray-400 font-medium">Data</th>
              <th className="text-right py-3 px-4 text-gray-400 font-medium">Ações</th>
            </tr>
          </thead>
          <tbody>
            <tr className="border-b border-gray-800">
              <td className="py-3 px-4 text-white">INV-001</td>
              <td className="py-3 px-4 text-white">R$ 100,00</td>
              <td className="py-3 px-4">
                <StatusBadge status="approved" />
              </td>
              <td className="py-3 px-4 text-gray-400">
                {new Date().toLocaleDateString()}
              </td>
              <td className="py-3 px-4 text-right">
                <div className="flex justify-end gap-2">
                  <Button
                    variant="ghost"
                    size="icon"
                    className="text-gray-400 hover:text-white"
                    asChild
                  >
                    <Link href="/invoices/INV-001">
                      <Eye className="h-4 w-4" />
                    </Link>
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    className="text-gray-400 hover:text-white"
                  >
                    <Download className="h-4 w-4" />
                  </Button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}
