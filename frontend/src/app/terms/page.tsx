import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "利用規約 | Menta Clone",
  description: "Menta Clone の利用規約",
};

export default function TermsPage() {
  return (
    <div className="flex flex-col min-h-screen">
      <header className="border-b bg-background/80 backdrop-blur-sm sticky top-0 z-50">
        <div className="max-w-6xl mx-auto px-6 h-16 flex items-center">
          <a href="/" className="text-xl font-bold tracking-tight">
            Menta Clone
          </a>
        </div>
      </header>
      <main className="flex-1 py-16 px-6">
        <div className="max-w-3xl mx-auto">
          <h1 className="text-3xl font-bold mb-8">利用規約</h1>
          <p className="text-sm text-muted-foreground mb-8">
            最終更新日：2026年4月13日
          </p>

          <section className="mb-8">
            <h2 className="text-xl font-semibold mb-3">第1条（適用）</h2>
            <p className="text-sm text-muted-foreground leading-relaxed">
              本規約は、Menta Clone（以下「本サービス」）が提供するすべてのサービスの利用条件を定めるものです。ユーザーは本規約に同意した上で本サービスを利用するものとします。
            </p>
          </section>

          <section className="mb-8">
            <h2 className="text-xl font-semibold mb-3">第2条（利用登録）</h2>
            <p className="text-sm text-muted-foreground leading-relaxed">
              利用登録を申請した者が本規約に同意し、所定の方法により申請した場合に限り、本サービスの利用登録が完了するものとします。本サービスは、利用登録の申請者に以下の事由があると判断した場合、申請を承認しないことがあります。
            </p>
            <ul className="list-disc list-inside mt-3 text-sm text-muted-foreground space-y-1">
              <li>虚偽の事項を届け出た場合</li>
              <li>本規約に違反したことがある者からの申請である場合</li>
              <li>その他、本サービスが不適切と判断した場合</li>
            </ul>
          </section>

          <section className="mb-8">
            <h2 className="text-xl font-semibold mb-3">第3条（禁止事項）</h2>
            <p className="text-sm text-muted-foreground leading-relaxed">
              ユーザーは、本サービスの利用にあたり、以下の行為をしてはなりません。
            </p>
            <ul className="list-disc list-inside mt-3 text-sm text-muted-foreground space-y-1">
              <li>法令または公序良俗に違反する行為</li>
              <li>犯罪行為に関連する行為</li>
              <li>本サービスのサーバーまたはネットワークの機能を破壊、妨害する行為</li>
              <li>他のユーザーに関する個人情報等を収集または蓄積する行為</li>
              <li>他のユーザーに成りすます行為</li>
              <li>その他、本サービスが不適切と判断する行為</li>
            </ul>
          </section>

          <section className="mb-8">
            <h2 className="text-xl font-semibold mb-3">第4条（免責事項）</h2>
            <p className="text-sm text-muted-foreground leading-relaxed">
              本サービスは、本サービスに起因してユーザーに生じたあらゆる損害について、本サービスの故意または重大な過失による場合を除き、一切の責任を負いません。
            </p>
          </section>

          <section className="mb-8">
            <h2 className="text-xl font-semibold mb-3">第5条（規約の変更）</h2>
            <p className="text-sm text-muted-foreground leading-relaxed">
              本サービスは、必要と判断した場合には、ユーザーへの通知なく本規約を変更することができるものとします。変更後の規約は、本サービス上に掲示した時点から効力を生じるものとします。
            </p>
          </section>

          <section className="mb-8">
            <h2 className="text-xl font-semibold mb-3">第6条（準拠法・管轄裁判所）</h2>
            <p className="text-sm text-muted-foreground leading-relaxed">
              本規約の解釈にあたっては、日本法を準拠法とします。本サービスに関して紛争が生じた場合には、東京地方裁判所を第一審の専属的合意管轄裁判所とします。
            </p>
          </section>
        </div>
      </main>
      <footer className="border-t py-8 px-6 text-center text-sm text-muted-foreground">
        © 2026 Menta Clone. All rights reserved.
      </footer>
    </div>
  );
}
