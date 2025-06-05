import { Router, RouterModule } from '@angular/router';
import { UserService } from './../../../services/user.service';
import { Component, inject, OnInit, TemplateRef } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { AsyncPipe, CommonModule } from '@angular/common';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgbHighlight } from '@ng-bootstrap/ng-bootstrap';
import { map, Observable, startWith } from 'rxjs';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import {  FormDirective, InputGroupComponent, InputGroupTextDirective } from '@coreui/angular';
import { IconDirective } from '@coreui/icons-angular';

interface ShopData {
  id: string;
  name: string;
  role: string;
  isActive: boolean;
}

@Component({
  selector: 'app-main',
  standalone: true,
  imports: [AsyncPipe, ReactiveFormsModule, NgbHighlight, CommonModule, RouterModule,FormsModule,
      FormDirective,
      InputGroupComponent,
      InputGroupTextDirective,
      IconDirective,
    ],
  templateUrl: './main.component.html',
})
export class MainComponent implements OnInit {
  // Define the observable for shop data
  shopName = ''
  shopDomain = ''

  shopData$: Observable<ShopData[]>;
  filter = new FormControl('', { nonNullable: true });
  shopData: ShopData[] = [];
  private modalService = inject(NgbModal);
  
  constructor(
    private readonly userService: UserService,
    private readonly router: Router,
    private readonly toast: ToastrService
  ) {
    this.shopData$ = this.filter.valueChanges.pipe(
            startWith(''),
            map((text) => this.search(text))
          );
  }

  search(text: string): ShopData[] {
    return this.shopData.filter((shop) => {
      const term = text.toLowerCase();
      return shop.name.toLowerCase().includes(term);
    });
  }

  getToken() {
    const authData = localStorage.getItem(this.userService.authKey);
    if (!authData) {
      this.toast.error('You are not logged in. Please login first.', 'Error');
      this.router.navigate(['/login']);
      return;
    }
    const parsedAuthData = JSON.parse(authData);
    if (!parsedAuthData || !parsedAuthData.token) {
      this.toast.error(
        'Invalid authentication data. Please login again.',
        'Error'
      );
      this.router.navigate(['/login']);
      return;
    }
    return parsedAuthData.token;
  }

  ngOnInit(): void {
    const token = this.getToken();
    this.userService.getAllShops(token).subscribe({
      next: (response: any) => {
        if (response && response.data && Array.isArray(response.data)) {
          this.shopData = response.data.map((shop: any) => ({
            id: shop.id,
            name: shop.name,
            role: shop.role || 'UNKNOWN',
            isActive: shop.is_active || false,
          }));
          this.shopData$ = this.filter.valueChanges.pipe(
            startWith(''),
            map((text) => this.search(text))
          );
          if (this.shopData.length === 0) {
            this.toast.warning(
              'You have no shops. Please create a shop first.',
              'Warning'
            );
          }
        } else {
          this.toast.error(
            'Failed to load shops. Please try again later.',
            'Error'
          );
          console.error('Invalid response format:', response);
          this.shopData = [];
        }
      },
      error: (error) => {
        this.toast.error(
          'Failed to load shops. Please try again later.',
          'Error'
        );
        console.error('Error fetching shops:', error);
      },
    });
  }

  startWorkOnShop(shopId: string) {
    this.router.navigate(['/shop', shopId, 'work']);
  }

  goToShop(shopId: string) {
    this.router.navigate(['/shop', shopId]);
  }

  leaveShop(shopId: string) {
    throw new Error(
      `Method not implemented. That means you are unable to leave shop with id: ${shopId}`
    );
  }

  createNewShop() {
   const token = this.getToken();
   if (this.shopDomain.length < 3 ||this.shopName.length < 3) {
     this.toast.error('Shop name and domain must be at least 3 characters long.', 'Error');
     return;
   }

    this.userService.createNewShop(this.shopName, this.shopDomain, token).subscribe({
      next: (response: any) => {
        if (response) {
          this.toast.success('Shop created successfully!', 'Success');
          // refresh page to show new shop
          window.location.reload();
          this.modalService.dismissAll();
        } else {
          this.toast.error('Failed to create shop. Please try again.', 'Error');
        }
      },
      error: (error) => {
        console.error('Error creating shop:', error);
        this.toast.error('Failed to create shop. Please try again.', 'Error');
      },
    });
  }

  	openModal(content: TemplateRef<any>) {
		this.modalService.open(content, { centered: true });
	}

}
